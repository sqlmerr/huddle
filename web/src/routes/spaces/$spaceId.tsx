import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/spaces/$spaceId')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/spaces/$spaceId"!</div>
}
