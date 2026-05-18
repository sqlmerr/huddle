import { useAuth } from '#/context/auth'
import { loginRequestSchema, type LoginRequest } from '#/lib/schemas'
import { createFileRoute, useNavigate, Link } from '@tanstack/react-router'
import { useForm } from '@tanstack/react-form'
import { toast } from 'sonner'
import { ApiError, login } from '#/lib/api'
import { HuddleLogo } from '#/components/HuddleLogo'
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '#/components/ui/card'
import { Label } from '#/components/ui/label'
import { Input } from '#/components/ui/input'
import { Button } from '#/components/ui/button'
import { useMutation } from '@tanstack/react-query'

export const Route = createFileRoute('/login')({
  component: LoginPage,
})

function LoginPage() {
  const navigate = useNavigate()
  const { login: authLogin } = useAuth()

  const mutation = useMutation({
    mutationFn: (data: LoginRequest) => {
      return login(data)
    },
    onSuccess: (data) => {
      authLogin(data.accessToken)
      toast.success('Welcome back!')
      navigate({ to: '/' }) // TODO: /dashboard
    },
    onError: (error) => {
      console.log(error)
      toast.error(
        error instanceof ApiError ? JSON.stringify(error) : 'Login failed',
      )
    },
  })

  const form = useForm({
    defaultValues: {
      username: '',
      password: '',
    },
    onSubmit({ value }) {
      mutation.mutate(loginRequestSchema.parse(value))
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
            <CardTitle className="text-2xl">Welcome back</CardTitle>
            <CardDescription>
              Enter your credentials to access your workspace
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
              {"Don't have an account? "}
              <Link to="/register" className="text-primary hover:underline">
                Sign up
              </Link>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
