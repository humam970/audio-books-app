import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
    return twMerge(clsx(inputs));
}

// scroll based
// view port detection
// sticky position
// easing
// text splitting (gsap splitText, splitType)
// math.map
// lerp
// request animation frame
