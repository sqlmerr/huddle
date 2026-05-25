import type { Space } from '#/lib/schemas'
import { Link } from '@tanstack/react-router'
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from './ui/card'
import { Folder, User } from 'lucide-react'

export function SpaceCard({ space }: { space: Space }) {
  return (
    <Link to={`/spaces/${space.id}`}>
      <Card className="h-full transition-colors hover:border-primary/50 hover:bg-surface/80">
        <CardHeader>
          <div className="flex items-start justify-between">
            <div className="rounded-lg bg-primary/10 p-2">
              <Folder className="h-5 w-5 text-primary" />
            </div>
          </div>
          <CardTitle className="line-clamp-1 text-lg">{space.title}</CardTitle>
          {space.description && (
            <CardDescription className="line-clamp-2">
              {space.description}
            </CardDescription>
          )}
        </CardHeader>
        <CardContent>
          <div className="flex items-center gap-2 text-sm text-muted">
            <User className="h-4 w-4" />
            <span>{space.ownerId}</span>
          </div>
        </CardContent>
      </Card>
    </Link>
  )
}
