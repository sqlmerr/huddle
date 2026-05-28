import { useState } from 'react'
import { ChevronLeft, ChevronRight, Layout, Plus } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { ScrollArea } from '@/components/ui/scroll-area'
import { HuddleLogo } from '@/components/HuddleLogo'
import { cn } from '@/lib/utils'
import { Link, useParams } from '@tanstack/react-router'
import { useSpaceContext } from '#/context/space'

interface SidebarProps {
  onAddBoard: () => void
}

export function Sidebar({ onAddBoard }: SidebarProps) {
  const [collapsed, setCollapsed] = useState(false)
  const { space, boards } = useSpaceContext()
  const { boardId } = useParams({ strict: false })

  return (
    <aside
      className={cn(
        'flex flex-col border-r border-border bg-surface transition-all duration-300',
        collapsed ? 'w-16' : 'w-64',
      )}
    >
      {/* Logo */}
      <div className="flex h-16 items-center justify-between border-b border-border px-4">
        {collapsed ? (
          <Link
            to="/dashboard"
            className="flex h-9 w-9 items-center justify-center rounded-lg bg-primary text-sm font-bold text-foreground"
          >
            H
          </Link>
        ) : (
          <Link to="/dashboard">
            <HuddleLogo size="sm" />
          </Link>
        )}
        <Button
          variant="ghost"
          size="icon"
          className="h-8 w-8"
          onClick={() => setCollapsed(!collapsed)}
        >
          {collapsed ? (
            <ChevronRight className="h-4 w-4" />
          ) : (
            <ChevronLeft className="h-4 w-4" />
          )}
        </Button>
      </div>

      {/* Space name */}
      {!collapsed && (
        <div className="border-b border-border px-4 py-3">
          <h2 className="truncate font-semibold text-foreground">
            {space.title}
          </h2>
          {space.description && (
            <p className="truncate text-xs text-muted">{space.description}</p>
          )}
        </div>
      )}

      {/* Boards list */}
      <ScrollArea className="flex-1">
        <div className={cn('py-2', collapsed ? 'px-2' : 'px-3')}>
          {!collapsed && (
            <p className="mb-2 px-2 text-xs font-medium uppercase text-muted">
              Boards
            </p>
          )}
          <div className="space-y-1">
            {boards.map((board) => (
              <Link
                key={board.id}
                to={`/spaces/${space.id}/boards/${board.id}`}
                className={cn(
                  'flex items-center gap-2 rounded-md px-2 py-2 text-sm transition-colors hover:bg-border',
                  boardId === board.id && 'bg-border text-foreground',
                  collapsed && 'justify-center',
                )}
              >
                <Layout className="h-4 w-4 shrink-0 text-primary" />
                {!collapsed && <span className="truncate">{board.title}</span>}
              </Link>
            ))}
          </div>
        </div>
      </ScrollArea>

      {/* Add board button */}
      <div className={cn('border-t border-border p-3', collapsed && 'px-2')}>
        <Button
          variant="outline"
          className={cn('w-full', collapsed && 'px-0')}
          onClick={onAddBoard}
        >
          <Plus className="h-4 w-4" />
          {!collapsed && <span className="ml-2">Add board</span>}
        </Button>
      </div>
    </aside>
  )
}
