import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/genres')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/genres"!</div>
}
