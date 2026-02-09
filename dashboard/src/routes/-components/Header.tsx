import { BellIcon, HeadphonesIcon, LibraryIcon, UserIcon } from "lucide-react";
import Navbar from "./Navbar";

function Header() {
    return (
        <header
            className="
            flex justify-between items-center
            bg-card text-card-foreground
            shadow-sm px-gutter
            text-base
            h-18
            "
        >
            <a href="#" className="text-primary transition-colors hover:opacity-80">
                <HeadphonesIcon size={44} />
            </a>

            <Navbar />

            <div className="flex items-center gap-4 text-primary">
                <button className="p-2 rounded-full hover:bg-accent hover:text-accent-foreground transition-colors">
                    <BellIcon size={24} />
                </button>
                <button className="p-2 rounded-full hover:bg-accent hover:text-accent-foreground transition-colors">
                    <LibraryIcon size={24} />
                </button>
                <button className="p-2 rounded-full hover:bg-accent hover:text-accent-foreground transition-colors">
                    <UserIcon size={24} />
                </button>
            </div>
        </header>
    );
}

export default Header;
