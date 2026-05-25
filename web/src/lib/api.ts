import z from 'zod'
import {
  loginResponseSchema,
  errorSchema,
  type LoginRequest,
  type LoginResponse,
  type User,
  type ErrorResponse,
  type RegisterRequest,
  type RegisterResponse,
  registerResponseSchema,
  type GetMySpacesResponse,
  getMySpacesResponseSchema,
  type CreateSpaceRequest,
  type CreateSpaceResponse,
  createSpaceResponseSchema,
  type GetMeResponse,
  getMeResponseSchema,
} from './schemas'

const BASE_URL = import.meta.env.VITE_API_URL ?? 'http://localhost:8080'

export class ApiError extends Error {
  public readonly error: string
  public readonly statusCode: number

  constructor(error: string, message: string, statusCode: number) {
    super(message)
    this.name = 'ApiError'
    this.error = error
    this.statusCode = statusCode

    // Maintains proper stack trace for where our error was thrown (only available on V8)
    Error.captureStackTrace(this, ApiError)
  }
}

function toSnakeCase(obj: Record<string, unknown>): Record<string, unknown> {
  return Object.fromEntries(
    Object.entries(obj).map(([k, v]) => [
      k.replace(/[A-Z]/g, (c) => `_${c.toLowerCase()}`),
      v,
    ]),
  )
}

function toCamelCase(obj: unknown): unknown {
  if (Array.isArray(obj)) {
    return obj.map(toCamelCase)
  }

  if (obj !== null && typeof obj === 'object') {
    return Object.fromEntries(
      Object.entries(obj as Record<string, unknown>).map(([k, v]) => [
        k.replace(/_([a-z])/g, (_, c) => c.toUpperCase()),
        toCamelCase(v),
      ]),
    )
  }

  return obj
}

export async function apiFetch<T>(
  path: string,
  schema: z.ZodType<T>,
  options?: RequestInit,
): Promise<T> {
  const token = localStorage.getItem('token')
  const res = await fetch(`${BASE_URL}${path}`, {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
      ...options?.headers,
    },
  })

  if (!res.ok) {
    try {
      const errorJson = await res.json()
      console.log(errorJson)
      const errorData = errorSchema.parse(toCamelCase(errorJson))
      throw new ApiError(errorData.error, errorData.message, res.status)
    } catch (error) {
      if (error instanceof ApiError) {
        throw error
      }
      const errorText = await res.text().catch(() => '')
      throw new ApiError(
        'UNKNOWN_ERROR',
        errorText || `HTTP ${res.status}: ${res.statusText}`,
        res.status,
      )
    }
  }

  const json = await res.json()
  return schema.parse(toCamelCase(json))
}

export async function login(data: LoginRequest): Promise<LoginResponse> {
  return await apiFetch('/api/v1/auth/login', loginResponseSchema, {
    method: 'POST',
    body: JSON.stringify(toSnakeCase(data as Record<string, unknown>)),
  })
}

export async function register(
  data: RegisterRequest,
): Promise<RegisterResponse> {
  return await apiFetch('/api/v1/auth/register', registerResponseSchema, {
    method: 'POST',
    body: JSON.stringify(toSnakeCase(data as Record<string, unknown>)),
  })
}

export async function getMe(): Promise<GetMeResponse> {
  return await apiFetch('/api/v1/users/me', getMeResponseSchema, {
    method: 'GET',
  })
}

export async function getMySpaces(): Promise<GetMySpacesResponse> {
  const response = await apiFetch(
    '/api/v1/spaces/my',
    getMySpacesResponseSchema,
    {
      method: 'GET',
    },
  )

  console.log('r', response)
  return response
}

export async function createSpace(
  data: CreateSpaceRequest,
): Promise<CreateSpaceResponse> {
  return await apiFetch('/api/v1/spaces', createSpaceResponseSchema, {
    method: 'POST',
    body: JSON.stringify(toSnakeCase(data as Record<string, unknown>)),
  })
}
