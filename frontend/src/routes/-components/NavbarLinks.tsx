import { cn } from "@/lib/utils";
import type { FileRoutesByFullPath } from "@/routeTree.gen";
import { Link } from "@tanstack/react-router";

const Links: { name: string; to: keyof FileRoutesByFullPath }[] = [
    { name: "Browse", to: "/" },
    { name: "Library", to: "/library" },
    { name: "Genres", to: "/genres" },
    { name: "New Release", to: "/new_release" },
];

function NavbarLinks({ vertical = false }: { vertical?: boolean }) {
    return (
        <ul className={cn("flex gap-5", vertical ? "flex-col" : "")}>
            {Links.map((link) => (
                <li key={link.to}>
                    <Link
                        to={link.to as string}
                        inactiveProps={{ className: "text-muted-foreground" }}
                        activeProps={{ className: "text-primary font-bold" }}
                    >
                        {link.name}
                    </Link>
                </li>
            ))}
        </ul>
    );
}

export default NavbarLinks;
