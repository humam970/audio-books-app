import { ArrowLeftIcon, BellIcon, EllipsisIcon, UserIcon } from "lucide-react";
import FloatButton from "./ui/float_button";

export function UserAndNotificationFloat() {
	return (
		<div className="flex justify-between mt-5">
			<FloatButton>
				<UserIcon color="white" />
			</FloatButton>
			<FloatButton>
				<BellIcon color="white" />
			</FloatButton>
		</div>
	);
}

export function BackAndOptionsFloat() {
	return (
		<div className="flex justify-between mt-5">
			<FloatButton>
				<ArrowLeftIcon color="white" />
			</FloatButton>
			<FloatButton>
				<EllipsisIcon color="white" />
			</FloatButton>
		</div>
	);
}
