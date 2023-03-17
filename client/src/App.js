import './App.css';
import {RouterProvider} from "react-router";
import {createBrowserRouter} from "react-router-dom";
import Home from "./Home";
import Login from "./Login";
import LoginCallback from "./LoginCallback";
import React, {useState} from "react";

const router = createBrowserRouter([
    {
        path: "/",
        element: <Home/>
    },
    {
        path: "/login",
        element: <Login/>
    },
    {
        path: "/callback",
        element: <LoginCallback/>
    }
])

export const AuthContext = React.createContext({
    user: null,
    setUser: (user) => {
    },
});

function App() {
    const [user, setUser] = useState(null);

    return (
        <div className="App">
            <AuthContext.Provider value={{user, setUser}}>
                <RouterProvider router={router}/>
            </AuthContext.Provider>
        </div>
    );
}

export default App;
