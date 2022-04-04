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
    FindAll: (page= 1, limit= 10) => {
        let {data, error, mutate} = useSWR(`/api/v1/clients?page=${page}&limit=${limit}`);
        let clients;
        if (!data) {
            clients = [];
            return {clients, error, mutate};
        }
        if (data) {
            data = data.data;
            return {data, error, mutate};
        }
        return {data, error, mutate};
    },
}
