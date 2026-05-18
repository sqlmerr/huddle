import { Outlet, createRootRoute } from '@tanstack/react-router'
import { TanStackRouterDevtoolsPanel } from '@tanstack/react-router-devtools'
import { TanStackDevtools } from '@tanstack/react-devtools'

import '../styles.css'
import { ThemeProvider } from '#/components/theme-provider'
import { AuthProvider } from '#/context/auth'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { Toaster } from '#/components/ui/sonner'

export const Route = createRootRoute({
  component: RootComponent,
})

const queryClient = new QueryClient()

function RootComponent() {
  return (
    <>
      <QueryClientProvider client={queryClient}>
        <ThemeProvider>
          <AuthProvider>
            <Outlet />
            <Toaster />
          </AuthProvider>
        </ThemeProvider>
      </QueryClientProvider>
      <TanStackDevtools
        config={{
          position: 'bottom-right',
        }}
        plugins={[
          {
            name: 'TanStack Router',
            render: <TanStackRouterDevtoolsPanel />,
          },
        ]}
      />
    </>
  )
}
