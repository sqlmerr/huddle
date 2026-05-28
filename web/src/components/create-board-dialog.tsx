import { ApiError, createBoard } from '#/lib/api'
import type { CreateBoardRequest } from '#/lib/schemas'
import { useForm } from '@tanstack/react-form'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { toast } from 'sonner'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from './ui/dialog'
import { Button } from './ui/button'
import { Label } from './ui/label'
import { Input } from './ui/input'

interface CreateBoardDialogProps {
  spaceId: string
  open: boolean
  onOpenChange: (open: boolean) => void
}

export function CreateBoardDialog({
  spaceId,
  open,
  onOpenChange,
}: CreateBoardDialogProps) {
  const queryClient = useQueryClient()
  const mutation = useMutation({
    mutationFn: (data: CreateBoardRequest) => {
      return createBoard(data)
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['space', spaceId, 'boards'] })
      toast.success('Successfully created new board')
      onOpenChange(false)
    },
    onError: (error) => {
      toast.error(
        error instanceof ApiError
          ? JSON.stringify(error)
          : 'Failed to create space',
      )
    },
  })

  const form = useForm({
    defaultValues: {
      title: '',
    },
    onSubmit({ value }) {
      mutation.mutate({ title: value.title, spaceId })
    },
  })

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Create a new board</DialogTitle>
          <DialogDescription>
            Boards contain lists and tasks for organizing your work.
          </DialogDescription>
        </DialogHeader>
        <form
          onSubmit={(e) => {
            e.preventDefault()
            form.handleSubmit()
          }}
          className="space-y-4"
        >
          <form.Field
            name="title"
            children={(field) => (
              <div className="space-y-2">
                <Label htmlFor="board-title">Title</Label>
                <Input
                  id="board-title"
                  placeholder="e.g., Sprint 1"
                  onBlur={field.handleBlur}
                  onChange={(e) => field.handleChange(e.target.value)}
                />
                <p className="text-sm text-red-500">
                  {field.state.meta.errors.join(', ')}
                </p>
              </div>
            )}
          />
          <DialogFooter>
            <div className="flex justify-end gap-2">
              <Button
                type="button"
                variant="outline"
                onClick={() => onOpenChange(false)}
              >
                Cancel
              </Button>
              <Button type="submit" disabled={mutation.isPending}>
                {mutation.isPending ? 'Creating...' : 'Create Board'}
              </Button>
            </div>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  )
}
