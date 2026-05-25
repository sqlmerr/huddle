import { CreateSpaceDialog } from '#/components/create-space-dialog'
import { Header } from '#/components/header'
import { SpaceCard } from '#/components/space-card'
import { useAuth } from '#/context/auth'
import { ApiError, createSpace, getMySpaces } from '#/lib/api'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { createFileRoute, useNavigate } from '@tanstack/react-router'
import { Folder } from 'lucide-react'
import { toast } from 'sonner'

export const Route = createFileRoute('/dashboard')({
  component: RouteComponent,
})

function RouteComponent() {
  const navigate = useNavigate()
  const { isAuthorized, logout, isLoading } = useAuth()
  const query = useQuery({ queryKey: ['spaces'], queryFn: getMySpaces })

  if (isLoading) {
    return (
      <div className="flex min-h-screen items-center justify-center">
        <div className="text-muted">Loading...</div>
      </div>
    )
  }

  if (!isAuthorized) {
    toast.error("Can't access this page")
    logout()
    navigate({ to: '/login' })

    return null
  }

  console.log(query.data)
  return (
    <div className="min-h-screen bg-background">
      <Header />

      <main className="mx-auto max-w-7xl px-6 py-8">
        <div className="mb-8 flex items-center justify-between">
          <div>
            <h1 className="text-2xl font-bold text-foreground">Your Spaces</h1>
            <p className="text-muted">Select a space to view its boards</p>
          </div>
          <CreateSpaceDialog />
        </div>

        {query.isLoading ? (
          <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
            {[...Array(3)].map((_, i) => (
              <div
                key={i}
                className="h-48 animate-pulse rounded-lg border border-border bg-surface"
              />
            ))}
          </div>
        ) : query.data?.data.length === 0 ? (
          <div className="flex flex-col items-center justify-center rounded-lg border border-dashed border-border py-16 text-center">
            <div className="mb-4 rounded-full bg-surface p-4">
              <Folder className="h-8 w-8 text-muted" />
            </div>
            <h2 className="mb-2 text-lg font-medium text-foreground">
              No spaces yet
            </h2>
            <p className="mb-6 max-w-sm text-muted">
              Create your first space to start organizing your work
            </p>
            <CreateSpaceDialog />
          </div>
        ) : (
          <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
            {query.data?.data.map((space) => (
              <SpaceCard key={space.id} space={space} />
            ))}
          </div>
        )}
      </main>
    </div>
  )
}
