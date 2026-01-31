import { SidebarProvider, SidebarContent, SidebarMenuButton, Sidebar } from "@/components/ui/sidebar";
import { Link } from "@tanstack/react-router";

function AppSidebar() {
    return (
        <SidebarProvider>
            <Sidebar collapsible="none">
                <SidebarContent
                    className="
                    [&_.active]:text-blue-500
                    [&_.active]:font-bold
                    "
                >
                    <SidebarMenuButton asChild>
                        <Link to="/authors">Authors</Link>
                    </SidebarMenuButton>
                    <SidebarMenuButton asChild>
                        <Link to="/narrators">narrators</Link>
                    </SidebarMenuButton>
                    <SidebarMenuButton asChild>
                        <Link to="/genres">genres</Link>
                    </SidebarMenuButton>
                    <SidebarMenuButton asChild>
                        <Link to="/books">books</Link>
                    </SidebarMenuButton>
                    <SidebarMenuButton asChild>
                        <Link to="/chapters">chapters</Link>
                    </SidebarMenuButton>
                </SidebarContent>
            </Sidebar>
        </SidebarProvider>
    );
}

export default AppSidebar;
