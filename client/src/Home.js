import React, {useContext} from 'react'
import {Navigate} from "react-router";
import {AuthContext} from "./App";

export default function Home() {
    const { user } = useContext(AuthContext)
    if (!user) {
        return <Navigate to="/login" />
    }

    console.log(user)

    return  (
        <div>
            <h1>User Information</h1>
            <p>Name: {user.name}</p>
            <p>Email: {user.email}</p>
        </div>
    )
}