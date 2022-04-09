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
        // Just return if page or limit is undefined.
        if (page == undefined || limit == undefined) {return}

        // Datagrid or Table in NextJS starts with 0 (zero) based slide index.
        // So, we need to increase page number.
        page +=1;

        try {
            return await axios.get(`/api/v1/clients?page=${page}&limit=${limit}`)
        } catch (error) {
            return error.response;
        }
    },
}
