const links = ["one", "two"];

function Navigation() {
	return (
		<nav>
			{links.map((link) => (
				<li>
					<a href="">{link}</a>
				</li>
			))}
		</nav>
	);
}

export default Navigation;
