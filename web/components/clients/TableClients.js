import React from "react";
import axios from "axios";
import { ClientAPI } from "../../lib/api/client";
import { DataGrid } from "@mui/x-data-grid";

export default function TableClients() {
    const columns = [
        { field: 'id', headerName: 'ID', width: 10 },
        { field: 'name', headerName: 'Name', minWidth: 110 },
        { field: 'description', headerName: 'Description', flex: 1 },
    ]

    const [page, setPage] = React.useState(0);
    const [limit, setLimit] = React.useState(5);
    const [total, setTotal] = React.useState(0);
    const [rows, setRows] = React.useState([]);

    const [loading, setLoading] = React.useState(false);
    const [selectionModel, setSelectionModel] = React.useState([]);
    const prevSelectionModel = React.useRef(selectionModel);

    React.useEffect(() => {
        setLoading(true);

        const getClients = async () => {
            page +=1;
            const result = await axios.get(`/api/v1/clients?page=${page}&limit=${limit}`)
            const body = await result.data;
            const data = body.data;
    
            setRows(data.clients);
            setTotal(data.total);
        }

        getClients();
        setLoading(false);

        setTimeout(() => {
            setSelectionModel(prevSelectionModel.current);
        });

    }, [page]);

    return (
        <div style={{ height: 400, width: '100%' }}>
            <DataGrid
                rows={rows}
                columns={columns}
                pagination
                pageSize={limit}
                rowsPerPageOptions={[5]}
                rowCount={total}
                paginationMode="server"

                onPageChange={(newPage) => {
                    prevSelectionModel.current = selectionModel;
                    setPage(newPage);
                }}
                onSelectionModelChange={(newSelectionModel) => {
                    setSelectionModel(newSelectionModel);
                }}

                selectionModel={selectionModel}
                loading={loading}
            />
        </div>
    );
}
