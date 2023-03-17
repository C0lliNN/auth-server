import React, {useContext, useEffect} from 'react'
import {useSearchParams} from "react-router-dom";
import axios from "axios";
import {AuthContext} from "./App";
import {Navigate} from "react-router";

export default function LoginCallback() {
    const [searchParams, _] = useSearchParams();
    const {user, setUser} = useContext(AuthContext)
    const code = searchParams.get("code")
    // validate state ideally

    useEffect(() => {
        let form = new FormData();
        form.append("grant_type", 'authorization_code')
        form.append("redirect_uri", 'http://localhost:3000/callback')
        form.append("client_id", 'web')
        form.append("client_secret", 'test')
        form.append("code", code)
        form.append("scope", 'openid email profile')

        axios.post('http://localhost:9998/oauth/token', form, {headers: {"Content-Type": "multipart/form-data"}})
            .then(response => response.data)
            .then(data => {
                return axios.get('http://localhost:9998/userinfo', {
                    headers: {
                        'Authorization': `Bearer ${data.access_token}`
                    }
                }).then(response => ({...response.data, ...data}))

            })
            .then(user => setUser(user))
            .catch(err => console.error(err))

    }, [searchParams])

    if (user) {
        return <Navigate to="/"/>
    }

    return <div>Loading</div>
}