import React from 'react'

function serialize(obj) {
    const str = [];
    for (const p in obj)
        if (obj.hasOwnProperty(p)) {
            str.push(encodeURIComponent(p) + "=" + encodeURIComponent(obj[p]));
        }
    return str.join("&");
}

export default function Login() {
    const query = {
        client_id: 'web',
        prompt: 'Welcome+back',
        redirect_uri: "http://localhost:3000/callback",
        response_type: 'code',
        scope: 'openid email profile',
        state: '5712db2d-3d0b-42d7-921c-26457296a704'
    }

    return (
        <div>
            <h1 style={{paddingBottom: "20px"}}>Authenticate</h1>
            <a className="login-btn" href={"http://localhost:9998/auth?" + serialize(query)}>Login</a>
        </div>
    )
}
