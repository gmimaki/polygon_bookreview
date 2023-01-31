import {createBrowserRouter, RouterProvider} from "react-router-dom";
import HelloWorld from './HelloWorld';

const router = createBrowserRouter([
	{
		path: "/",
		element: <HelloWorld/>,
	}
]);

const route = () => (
	<>
		<RouterProvider router={router}/>
	</>
);

export default route