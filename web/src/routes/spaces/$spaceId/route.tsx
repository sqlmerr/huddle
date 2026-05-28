import { CreateBoardDialog } from '#/components/create-board-dialog'
import { Sidebar } from '#/components/sidebar'
import { SpaceContextProvider } from '#/context/space'
import { getSpace, getSpaceBoards } from '#/lib/api'
import { useQuery } from '@tanstack/react-query'
import { createFileRoute, Outlet } from '@tanstack/react-router'
import { useState } from 'react'
import { toast } from 'sonner'

export const Route = createFileRoute('/spaces/$spaceId')({
  component: LayoutComponent,
})

function LayoutComponent() {
  const { spaceId } = Route.useParams()
  const navigate = Route.useNavigate()
  const [showCreateBoard, setShowCreateBoard] = useState(false)

  const spaceQuery = useQuery({
    queryKey: ['space', spaceId],
    queryFn: () => {
      return getSpace(spaceId)
    },
  })

  const boardsQuery = useQuery({
    queryKey: ['space', spaceId, 'boards'],
    queryFn: () => {
      return getSpaceBoards(spaceId)
    },
  })

  if (spaceQuery.isLoading || boardsQuery.isLoading || boardsQuery.isFetching) {
    return (
      <div className="flex h-screen items-center justify-center bg-background">
        <div className="h-8 w-8 animate-spin rounded-full border-2 border-primary border-t-transparent" />
      </div>
    )
  }

  if (spaceQuery.isError || !spaceQuery.data) {
    toast.error('Failed to fetch space data')
    console.error('Failed to fetch space data:', spaceQuery.error)
    navigate({ to: '/dashboard' })
    return null
  }

  if (boardsQuery.isError) {
    toast.error('Failed to fetch boards data')
    console.error('Failed to fetch boards data:', boardsQuery.error)
    navigate({ to: '/dashboard' })
    return null
  }

  return (
    <SpaceContextProvider
      space={spaceQuery.data}
      boards={boardsQuery.data?.data ?? []}
    >
      <div className="flex h-screen bg-background">
        <Sidebar onAddBoard={() => setShowCreateBoard(true)} />
        <main className="flex-1 overflow-hidden">
          {/* <Outlet context={{ space, boards }} />*/}
          <Outlet />
        </main>
        <CreateBoardDialog
          spaceId={spaceQuery.data.id}
          open={showCreateBoard}
          onOpenChange={setShowCreateBoard}
        />
      </div>
    </SpaceContextProvider>
  )
}
