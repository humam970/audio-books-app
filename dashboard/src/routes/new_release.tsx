import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/new_release')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/new_release"!</div>
}
