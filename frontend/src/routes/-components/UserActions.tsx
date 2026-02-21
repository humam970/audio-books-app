import { Button } from "@/components/ui/button";
import { BellIcon, LibraryIcon, UserIcon, type LucideIcon } from "lucide-react";

const Icons: LucideIcon[] = [BellIcon, LibraryIcon, UserIcon];

function UserActions() {
    return (
        <div className="flex items-center gap-4 text-primary">
            {Icons.map((Icon) => (
                <Button
                    asChild
                    size="icon"
                    variant="ghost"
                    className="p-2 rounded-full hover:bg-accent hover:text-accent-foreground transition-colors"
                >
                    <Icon />
                </Button>
            ))}
        </div>
    );
}

export default UserActions;
