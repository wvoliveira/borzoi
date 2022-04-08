import axios from "axios";
import React, {useEffect} from "react";
import useSWR from "swr";
import Router from "next/router";

export const ClientAPI = {
    Create: async(name, description) => {
        try {
            return await axios.post(
                '/api/v1/clients',
                JSON.stringify({"name": name, "description": description}), {
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
    FindAll: async(page, limit) => {
        if (page === undefined || page === 0) {
            page = 1
        }

        if (limit === undefined || limit === 0) {
            limit = 5
        }

        try {
            const res = await axios.get(
                `/api/v1/clients?page=${page}&limit=${limit}`,
                {
                    headers: {
                        "Accept": "application/json",
                        "Content-Type": "application/json",
                    },
                }
            );
            const data = await res.data;
            console.log(data);
            return { data: data };
        } catch (error) {
            return error.response;
        }
    },
}
