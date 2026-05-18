import { z } from 'zod'

export const errorSchema = z.object({
  error: z.string(),
  message: z.string(),
})

export const loginRequestSchema = z.object({
  username: z.string().min(3).max(32).optional(),
  email: z.email().optional(),
  password: z.string().min(1),
})

export const loginResponseSchema = z.object({
  accessToken: z.string(),
  tokenType: z.string(),
})

export const registerRequestSchema = z.object({
  username: z.string().min(3).max(32),
  email: z.email(),
  password: z.string().min(1),
})

export const userSchema = z.object({
  id: z.uuid(),
  username: z.string(),
  email: z.email(),
  createdAt: z.coerce.date(),
})

export const registerResponseSchema = userSchema

export type ErrorResponse = z.infer<typeof errorSchema>

// Auth
export type LoginRequest = z.infer<typeof loginRequestSchema>
export type LoginResponse = z.infer<typeof loginResponseSchema>
export type RegisterRequest = z.infer<typeof registerRequestSchema>
export type RegisterResponse = z.infer<typeof registerResponseSchema>

export type User = z.infer<typeof userSchema>
