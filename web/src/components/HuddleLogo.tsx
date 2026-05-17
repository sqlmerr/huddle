import { cn } from '@/lib/utils'

interface HuddleLogoProps {
  className?: string
  showText?: boolean
  size?: 'sm' | 'md' | 'lg'
}

export function HuddleLogo({
  className,
  showText = true,
  size = 'md',
}: HuddleLogoProps) {
  const sizeClasses = {
    sm: 'h-7 w-7 text-sm',
    md: 'h-9 w-9 text-base',
    lg: 'h-12 w-12 text-xl',
  }

  const textSizeClasses = {
    sm: 'text-lg',
    md: 'text-xl',
    lg: 'text-2xl',
  }

  return (
    <div className={cn('flex items-center gap-2', className)}>
      <div
        className={cn(
          'flex items-center justify-center rounded-lg bg-primary font-bold text-foreground',
          sizeClasses[size],
        )}
      >
        H
      </div>
      {showText && (
        <span
          className={cn('font-semibold text-foreground', textSizeClasses[size])}
        >
          Huddle
        </span>
      )}
    </div>
  )
}
