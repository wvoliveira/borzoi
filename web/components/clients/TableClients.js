import React from "react";
import {ClientAPI} from "../../lib/api/client";
import {DataGrid} from "@mui/x-data-grid";

export default function TableClients() {
    const [loading, setLoading] = React.useState(true);

    const {data, error, mutate} = ClientAPI.FindAll();

    const columns = [
        {field: 'id', headerName: 'ID', width: 10},
        {field: 'name', headerName: 'Name', minWidth: 110},
        {field: 'description', headerName: 'Description', flex: 1},
    ]

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);

        try {
            const {data, status} = ClientAPI.FindAll();
            if (status !== 200 && status !== 500) {
                //setError(data.message);
                console.log(data.message);
            }
            // setClients(data.data);
        } catch (error) {
            console.error(error);
        } finally {
            setLoading(false);
        }
    };

    const handlePageSize = async (pageSize) => {
        setLoading(true);
        try {
            const {data, status} = ClientAPI.FindAll({limit: pageSize});
            if (status !== 200 && status !== 500) {
                console.log(data.message);
            }
        } catch (error) {
            console.error(error);
        } finally {
            setLoading(false);
        }
    };

    const handlePage = async (page) => {
        setLoading(true);
        try {
            const {data, status} = ClientAPI.FindAll({page: page});
            if (status !== 200 && status !== 500) {
                console.log(data.message);
            }
        } catch (error) {
            console.error(error);
        } finally {
            setLoading(false);
        }
    };

    if (!data) {
        return (
            <div style={{height: 400, width: '80%'}}>
                <DataGrid
                    loading={true}
                    columns={columns}
                    rows={[]}/>
            </div>
        )
    }

    return (
        <div style={{height: 400, width: '80%'}}>
            <DataGrid
                rows={data.clients}
                columns={columns}
                pagination
                pageSize={data.per_page}
                rowsPerPageOptions={[10]}
                rowCount={data.total}
                paginationMode="server"
                onPageChange={handlePage}
                page={data.page}
            />
        </div>
    );
}
