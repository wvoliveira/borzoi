import axios from "axios";
import React from "react";

export const AuthAPI = {
    Check: async() => {
        try {
            return await axios.get(
                '/api/auth/check', {
                    headers: {
                        "Accept": "application/json",
                        "Content-Type": "application/json",
                    },
                }
            );
        } catch (error) {
            return error.response;
        }
    },
    Login: async(email, password) => {
        try {
            return await axios.post(
                '/api/auth/password/login',
                JSON.stringify({"email": email, "password": password}), {
                    headers: {
                        "Accept": "application/json",
                        "Content-Type": "application/json",
                    },
                }
            );
        } catch (error) {
            return error.response;
        }
    },
    Register: async(name, email, password) => {
        try {
            return await axios.post(
                '/api/auth/password/register',
                JSON.stringify({"name": name, "email": email, "password": password}), {
                    headers: {
                        "Content-Type": "application/json",
                    },
                }
            );
        } catch (error) {
            return error.response;
        }
    },
    Logout: async() => {
        try {
            return await axios.post(
                '/api/auth/logout',
                {
                    headers: {
                        "Content-Type": "application/json",
                    },
                }
            );
        } catch (error) {
            return error.response;
        }
    },
    Forgot: async(email) => {
        try {
            return await axios.post(
                '/api/auth/forgot',
                JSON.stringify({"email": email}), {
                    headers: {
                        "Content-Type": "application/json",
                    },
                }
            );
        } catch (error) {
            return error.response;
        }
    },
}
