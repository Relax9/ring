import {Outlet} from "react-router";

export default function Layout() {
    return (
        <div>
            <div>Layout Page</div>
            <Outlet/>
        </div>
    )
};