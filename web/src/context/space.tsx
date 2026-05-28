import type { Board, Space } from '#/lib/schemas'
import { createContext, useContext, type ReactNode } from 'react'

interface SpaceContextType {
  space: Space
  boards: Board[]
}

const SpaceContext = createContext<SpaceContextType | undefined>(undefined)

export function SpaceContextProvider({
  space,
  boards,
  children,
}: {
  space: Space
  boards: Board[]
  children: ReactNode
}) {
  return (
    <SpaceContext.Provider value={{ space, boards }}>
      {children}
    </SpaceContext.Provider>
  )
}

export function useSpaceContext() {
  const context = useContext(SpaceContext)
  if (context === undefined) {
    throw new Error('useSpaceContext must be used within SpaceContextProvider')
  }

  return context
}
