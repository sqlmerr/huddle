import { useSpaceContext } from '#/context/space'
import { createFileRoute } from '@tanstack/react-router'
import { Layout } from 'lucide-react'

export const Route = createFileRoute('/spaces/$spaceId/')({
  component: RouteComponent,
})

function RouteComponent() {
  const { space, boards } = useSpaceContext()

  return (
    <div className="flex h-full flex-col items-center justify-center p-8 text-center">
      <div className="mb-6 rounded-full bg-surface p-6">
        <Layout className="h-12 w-12 text-muted" />
      </div>
      <h1 className="mb-2 text-2xl font-bold text-foreground">{space.title}</h1>
      {space.description && (
        <p className="mb-6 max-w-md text-muted">{space.description}</p>
      )}
      {boards.length === 0 ? (
        <p className="text-muted">
          No boards yet. Create one from the sidebar to get started.
        </p>
      ) : (
        <p className="text-muted">
          Select a board from the sidebar to view tasks.
        </p>
      )}
    </div>
  )
}
