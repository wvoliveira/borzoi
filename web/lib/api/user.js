import axios from "axios";

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