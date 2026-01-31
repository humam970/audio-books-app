import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/narrators')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/narrators"!</div>
}
