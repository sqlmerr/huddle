import { AnimatedBackground } from '#/components/AnimatedBackground'
import { HuddleLogo } from '#/components/HuddleLogo'
import { ModeToggle } from '#/components/mode-toggle'
import { Button } from '#/components/ui/button'
import { createFileRoute, Link } from '@tanstack/react-router'
import { ArrowRight, Layout, Users, Zap } from 'lucide-react'

export const Route = createFileRoute('/')({ component: LandingPage })

function LandingPage() {
  return (
    <div className="relative min-h-screen bg-background">
      <AnimatedBackground />

      {/* Header */}
      <header className="relative z-10 flex items-center justify-between px-6 py-4 md:px-12">
        <HuddleLogo size="md" />
        <Button
          variant="outline"
          render={<Link to="/">Login</Link>}
          nativeButton={false}
        />
        {/* <ModeToggle /> */}
      </header>

      {/* Hero Section */}
      <main className="relative z-10 flex flex-col items-center justify-center px-6 py-24 text-center md:py-32 lg:py-40">
        <div className="mb-6 inline-flex items-center gap-2 rounded-full border border-border bg-surface/50 px-4 py-1.5 text-sm text-muted backdrop-blur-sm">
          <span className="h-2 w-2 rounded-full bg-primary animate-pulse" />
          Now in public beta
        </div>

        <h1 className="max-w-4xl text-balance text-4xl font-bold tracking-tight text-foreground sm:text-5xl md:text-6xl lg:text-7xl">
          Where teams get <span className="text-primary">things done</span>
        </h1>

        <p className="mt-6 max-w-2xl text-pretty text-lg text-muted md:text-xl">
          Huddle is a collaborative task manager that brings clarity to your
          workflow. Organize, prioritize, and ship together.
        </p>

        <div className="mt-10 flex flex-col gap-4 sm:flex-row">
          <Button
            size="lg"
            className="group"
            nativeButton={false}
            render={
              // Register
              <Link to="/">
                Get started free
                <ArrowRight className="h-4 w-4 transition-transform group-hover:translate-x-1" />
              </Link>
            }
          />

          <Button
            variant="outline"
            size="lg"
            // /Login
            render={<Link to="/">Sign in to your workspace</Link>}
            nativeButton={false}
          />
        </div>

        {/* Feature highlights */}
        <div className="mt-24 grid w-full max-w-4xl gap-6 sm:grid-cols-3">
          <FeatureCard
            icon={<Layout className="h-6 w-6" />}
            title="Kanban Boards"
            description="Visualize your workflow with intuitive drag-and-drop boards"
          />
          <FeatureCard
            icon={<Users className="h-6 w-6" />}
            title="Team Spaces"
            description="Organize projects into dedicated spaces for your team"
          />
          <FeatureCard
            icon={<Zap className="h-6 w-6" />}
            title="Lightning Fast"
            description="Built for speed with real-time updates and instant sync"
          />
        </div>
      </main>

      {/* Subtle footer text */}
      <div className="relative z-10 pb-8 text-center text-sm text-muted">
        Built for teams who ship
      </div>
    </div>
  )
}

interface FeatureCardProps {
  icon: React.ReactNode
  title: string
  description: string
}

function FeatureCard({ icon, title, description }: FeatureCardProps) {
  return (
    <div className="group rounded-xl border border-border bg-surface/50 p-6 backdrop-blur-sm transition-colors hover:border-primary/50 hover:bg-surface">
      <div className="mb-4 inline-flex rounded-lg bg-primary/10 p-3 text-primary">
        {icon}
      </div>
      <h3 className="mb-2 font-semibold text-foreground">{title}</h3>
      <p className="text-sm text-muted">{description}</p>
    </div>
  )
}
