import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/spaces/$spaceId/boards/$boardId')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/spaces/$spaceId/boards/$boardId"!</div>
}
