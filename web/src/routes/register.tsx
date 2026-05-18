import { HuddleLogo } from '#/components/HuddleLogo'
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '#/components/ui/card'
import { Input } from '#/components/ui/input'
import { Label } from '#/components/ui/label'
import { Button } from '#/components/ui/button'
import { ApiError, register } from '#/lib/api'
import { registerRequestSchema, type RegisterRequest } from '#/lib/schemas'
import { useForm } from '@tanstack/react-form'
import { useMutation } from '@tanstack/react-query'
import { createFileRoute, Link, useNavigate } from '@tanstack/react-router'
import { toast } from 'sonner'
import z, { ZodError } from 'zod'

export const Route = createFileRoute('/register')({
  component: RouteComponent,
})

function RouteComponent() {
  const navigate = useNavigate()

  const mutation = useMutation({
    mutationFn: (data: RegisterRequest) => {
      return register(data)
    },
    onSuccess: (data) => {
      toast.success('Account created successfully! Now login.')
      navigate({ to: '/login' })
    },
    onError: (error) => {
      console.log(error)
      toast.error(
        error instanceof ApiError ? JSON.stringify(error) : 'Register failed',
      )
    },
  })

  const form = useForm({
    defaultValues: {
      username: '',
      email: '',
      password: '',
    },
    onSubmit({ value }) {
      try {
        const data = registerRequestSchema.parse(value)
      } catch (error) {
        if (error instanceof ZodError) {
          console.log(z.prettifyError(error))
        }
        toast.error(
          error instanceof ZodError
            ? z.prettifyError(error)
            : 'Failed to parse data',
        )
      }
      mutation.mutate(registerRequestSchema.parse(value))
    },
  })

  return (
    <div className="flex min-h-screen items-center justify-center bg-background px-4">
      <div className="w-full max-w-md">
        <div className="mb-8 flex justify-center">
          <Link to="/">
            <HuddleLogo size="lg" />
          </Link>
        </div>

        <Card>
          <CardHeader className="text-center">
            <CardTitle className="text-2xl">Create an account</CardTitle>
            <CardDescription>
              Get started with your free workspace
            </CardDescription>
          </CardHeader>
          <CardContent>
            <form
              onSubmit={(e) => {
                e.preventDefault()
                e.stopPropagation()
                form.handleSubmit()
              }}
              className="space-y-4"
            >
              <form.Field
                name="username"
                children={(field) => (
                  <div className="space-y-2">
                    <Label htmlFor="username">Username</Label>
                    <Input
                      id="username"
                      placeholder="username"
                      onChange={(e) => field.handleChange(e.target.value)}
                      onBlur={field.handleBlur}
                      value={field.state.value}
                    />

                    <p className="text-sm text-red-500">
                      {field.state.meta.errors.join(', ')}
                    </p>
                  </div>
                )}
              />

              <form.Field
                name="email"
                children={(field) => (
                  <div className="space-y-2">
                    <Label htmlFor="email">Email</Label>
                    <Input
                      id="email"
                      placeholder="user@example.com"
                      onChange={(e) => field.handleChange(e.target.value)}
                      onBlur={field.handleBlur}
                      value={field.state.value}
                    />

                    <p className="text-sm text-red-500">
                      {field.state.meta.errors.join(', ')}
                    </p>
                  </div>
                )}
              />

              <form.Field
                name="password"
                children={(field) => (
                  <div className="space-y-2">
                    <Label htmlFor="password">Password</Label>
                    <Input
                      id="password"
                      type="password"
                      placeholder="password"
                      onChange={(e) => field.handleChange(e.target.value)}
                      onBlur={field.handleBlur}
                      value={field.state.value}
                    />
                    <p className="text-sm text-red-500">
                      {field.state.meta.errors.join(', ')}
                    </p>
                  </div>
                )}
              />

              <Button
                type="submit"
                className="w-full"
                disabled={mutation.isPending}
              >
                {mutation.isPending ? 'Signing in...' : 'Sign in'}
              </Button>

              {mutation.isError && (
                <p className="text-sm text-center text-destructive">
                  {mutation.error instanceof Error
                    ? mutation.error.message
                    : 'Login failed'}
                </p>
              )}
            </form>

            <div className="mt-6 text-center text-sm text-muted">
              Already have an account?{' '}
              <Link to="/login" className="text-primary hover:underline">
                Sign in
              </Link>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
