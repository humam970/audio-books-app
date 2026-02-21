import { Button } from "@/components/ui/button";
import { Sheet, SheetContent, SheetTrigger } from "@/components/ui/sheet";
import { MenuIcon } from "lucide-react";
import NavbarLinks from "./NavbarLinks";

function MobileNavbar() {
	return (
		<Sheet>
			<SheetTrigger asChild>
				<Button variant="ghost" size="icon">
					<MenuIcon size={24} />
				</Button>
			</SheetTrigger>

			<SheetContent side="left" className="w-[65vw] pt-4 ps-4">
				<nav>
					<NavbarLinks vertical />
				</nav>
			</SheetContent>
		</Sheet>
	);
}

export default MobileNavbar;
