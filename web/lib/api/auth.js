import axios from "axios";

export const AuthAPI = {
    Login: async(email, password) => {
        try {
            const response = await axios.post(
                '/api/v1/auth/password/login',
                JSON.stringify({ "email": email, "password": password }), {
                    headers: {
                        "Accept": "application/json",
                        "Content-Type": "application/json",
                    },
                }
            );
            return response;
        } catch (error) {
            return error.response;
        }
    },
    Register: async(name, email, password) => {
        try {
            const response = await axios.post(
                '/api/v1/auth/password/register',
                JSON.stringify({ user: { name, email, password } }), {
                    headers: {
                        "Content-Type": "application/json",
                    },
                }
            );
            return response;
        } catch (error) {
            return error.response;
        }
    },
}