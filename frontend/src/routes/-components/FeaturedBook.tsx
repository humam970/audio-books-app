import { Button } from "@/components/ui/button";
import { PlayIcon, DotIcon, StarIcon } from "lucide-react";

function FeaturedBook() {
    return (
        <section className="p-10 w-full rounded-2xl relative overflow-hidden isolate">
            <img
                src="/featured_book.png"
                className="absolute inset-0 w-full h-full object-cover object-center brightness-65 contrast-125 z-[-1]"
            />
            <div className="absolute inset-0 bg-linear-to-r from-blue-500/35 to-purple-500/35 z-[-1]"></div>

            <div className="mb-6 flex flex-col gap-3 items-start">
                <small className="text-xs text-white bg-white/20 border border-white/40 px-2 py-px rounded-full">
                    Featured This Week
                </small>
                <div className="flex flex-col gap-2 text-secondary">
                    <h2 className="text-4xl text-white font-bold">The Green Garden</h2>
                    <span>by Ali Abbas</span>
                    <div className="flex items-center gap-0 text-base">
                        <span>12h 45m</span>
                        <DotIcon />
                        <span>Fantasy</span>
                        <DotIcon />
                        <StarIcon size={14} className="fill-current" />
                        <span> 4.8</span>
                    </div>
                </div>

                <Button>
                    <PlayIcon size={24} />
                    Start Listening
                </Button>
            </div>
        </section>
    );
}

export default FeaturedBook;
