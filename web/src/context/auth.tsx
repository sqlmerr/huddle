import { getMe } from '#/lib/api'
import type { User } from '#/lib/schemas'
import { useQuery } from '@tanstack/react-query'
import { createContext, useContext, useState, useEffect } from 'react'
import type { ReactNode } from 'react'

interface AuthContextType {
  user: User | null
  token: string | null
  isLoading: boolean
  isAuthorized: boolean
  login: (token: string) => void
  logout: () => void
}

const AuthContext = createContext<AuthContextType | undefined>(undefined)

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<User | null>(null)
  const [token, setToken] = useState<string | null>(null)
  const [isLoading, setIsLoading] = useState(true)
  const [isAuthorized, setIsAuthorized] = useState(false)

  const query = useQuery({ queryKey: ['me'], queryFn: getMe, enabled: false })

  useEffect(() => {
    const storedToken = localStorage.getItem('token')
    if (storedToken) {
      setToken(storedToken)
      setIsAuthorized(true)
      query.refetch()
    }
    setIsLoading(false)
  }, [])

  useEffect(() => {
    if (query.isSuccess) {
      setUser(query.data)
    }
  }, [query.isSuccess, query.data])

  const login = (newToken: string) => {
    localStorage.setItem('token', newToken)

    const token = localStorage.getItem('token')
    setToken(newToken)
    query.refetch()
    setIsAuthorized(true)
  }

  const logout = () => {
    localStorage.removeItem('token')
    setToken(null)
    setUser(null)
  }

  return (
    <AuthContext.Provider
      value={{ user, token, isLoading, login, logout, isAuthorized }}
    >
      {children}
    </AuthContext.Provider>
  )
}

export function useAuth() {
  const context = useContext(AuthContext)
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider')
  }
  return context
}
