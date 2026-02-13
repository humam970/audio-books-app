import { Outlet, createRootRouteWithContext } from "@tanstack/react-router";
import { TanStackRouterDevtoolsPanel } from "@tanstack/react-router-devtools";
import { FormDevtoolsPanel } from "@tanstack/react-form-devtools";
import { TanStackDevtools } from "@tanstack/react-devtools";

import TanStackQueryDevtools from "../integrations/tanstack-query/devtools";

import type { QueryClient } from "@tanstack/react-query";
import { ThemeProvider } from "@/components/theme-provider";
import Header from "./-components/Header";

interface MyRouterContext {
    queryClient: QueryClient;
}

export const Route = createRootRouteWithContext<MyRouterContext>()({
    component: () => (
        <>
            <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
                <Header />
                <Outlet />
                <TanStackDevtools
                    config={{
                        position: "bottom-right",
                    }}
                    plugins={[
                        {
                            name: "Tanstack Router",
                            render: <TanStackRouterDevtoolsPanel />,
                        },
                        {
                            name: "Tanstack Form",
                            render: <FormDevtoolsPanel />,
                        },
                        TanStackQueryDevtools,
                    ]}
                />
            </ThemeProvider>
        </>
    ),
});
