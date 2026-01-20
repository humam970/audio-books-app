import type { ReactNode } from "react";

function FloatButton({
	children,
	className,
	...props
}: React.ComponentProps<"button"> & { children: ReactNode }) {
	return (
		<button
			{...props}
			className={
				className ??
				"size-12 bg-black rounded-full flex justify-center items-center"
			}
		>
			{children}
		</button>
	);
}

export default FloatButton;
