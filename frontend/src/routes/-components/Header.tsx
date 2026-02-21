import UserActions from "./UserActions";
import DesktopNavbar from "./DesktopNavbar";
import MobileNavbar from "./MobileNavbar";
import { useIsMobile } from "@/hooks/use-mobile";

function Header() {
    const isMobile = useIsMobile();

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
            {!isMobile ? <DesktopNavbar /> : <MobileNavbar />}

            <UserActions />
        </header>
    );
}

export default Header;
