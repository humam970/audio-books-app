import {
	// BackAndOptionsFloat,
	UserAndNotificationFloat,
} from "./../components/action_floats";

function Header() {
	return (
		<header className="col-[feature] *:mb-6">
			<UserAndNotificationFloat />
			<h1 className="text-2xl">
				Hey,
				<span className="text-[#E36166]">Steve!</span>
				<br />
				What will you listen today?
			</h1>
		</header>
	);
}

export default Header;
