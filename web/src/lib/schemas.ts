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

export const getMeResponseSchema = userSchema

export const spaceSchema = z.object({
  id: z.uuid(),
  title: z.string(),
  description: z.string().optional(),
  ownerId: z.uuid(),
  createdAt: z.coerce.date(),
  isArchived: z.boolean(),
})

export const getMySpacesResponseSchema = z.object({
  data: spaceSchema.array(),
})

export const createSpaceRequestSchema = z.object({
  title: z.string().min(1).max(50),
  description: z.string().min(1).max(1000).optional(),
})

export const createSpaceResponseSchema = spaceSchema

export type ErrorResponse = z.infer<typeof errorSchema>

// Auth
export type LoginRequest = z.infer<typeof loginRequestSchema>
export type LoginResponse = z.infer<typeof loginResponseSchema>
export type RegisterRequest = z.infer<typeof registerRequestSchema>
export type RegisterResponse = z.infer<typeof registerResponseSchema>
export type GetMeResponse = z.infer<typeof getMeResponseSchema>

// Spaces
export type GetMySpacesResponse = z.infer<typeof getMySpacesResponseSchema>
export type CreateSpaceRequest = z.infer<typeof createSpaceRequestSchema>
export type CreateSpaceResponse = z.infer<typeof createSpaceResponseSchema>

export type User = z.infer<typeof userSchema>
export type Space = z.infer<typeof spaceSchema>
