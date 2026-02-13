import type { FileRoutesByFullPath } from "@/routeTree.gen";
import { Link } from "@tanstack/react-router";

const Links: { name: string; to: keyof FileRoutesByFullPath }[] = [
    { name: "Browse", to: "/" },
    { name: "Library", to: "/library" },
    { name: "Genres", to: "/genres" },
    { name: "New Release", to: "/new_release" },
];

function Navbar() {
    return (
        <nav>
            <ul className="flex gap-9">
                {Links.map((link) => {
                    return (
                        <li
                            key={link.to}
                            className="text-muted-foreground hover:text-foreground transition-colors [&_.active]:text-primary [&_.active]:font-bold"
                        >
                            <Link to={link.to as string}>{link.name}</Link>
                        </li>
                    );
                })}
            </ul>
        </nav>
    );
}

export default Navbar;
