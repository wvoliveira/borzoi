import axios from "axios";

const fetcher = (...args) => fetch(...args).then((res) => res.json())

export const UserAPI = {
    Me: async() => {
        try {
            const response = await axios.get(
                '/api/v1/users/me', {
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
}
