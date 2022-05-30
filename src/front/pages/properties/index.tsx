import type { NextPage } from "next";
import { ApolloClient, gql, InMemoryCache } from "@apollo/client";
import { useEffect, useState } from "react";
import { CreateSection } from "components/CreateSection";
import { DeleteSection } from "components/DeleteSection";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { format } from "date-fns";
import ja from "date-fns/locale/ja";

const columns: GridColDef[] = [
  { field: "id", headerName: "ID", width: 50 },
  {
    field: "name",
    headerName: "物件名",
    width: 200,
    editable: true,
  },
  {
    field: "created_at",
    headerName: "作成日",
    description: "This column has a value getter and is not sortable.",
    sortable: false,
    width: 250,
    valueFormatter: (params) => {
      if (params.value) {
        const ts = Date.parse(params.value);
        const dt: Date = new Date(ts);
        return format(dt, "yyyy-MM-dd (EEEE) HH:mm:ss", { locale: ja });
      }
      return "";
    },
  },
  {
    field: "updated_at",
    headerName: "最終更新日",
    description: "This column has a value getter and is not sortable.",
    sortable: false,
    width: 250,
    valueFormatter: (params) => {
      if (params.value) {
        const ts = Date.parse(params.value);
        const dt: Date = new Date(ts);
        return format(dt, "yyyy-MM-dd (EEEE) HH:mm:ss", { locale: ja });
      }
      return "";
    },
  },
];

const Sections: NextPage = () => {
  const [rows, setRowData] = useState([]);
  const [isLoading, setIsLoading] = useState<boolean>(true);

  useEffect(() => {
    const client = new ApolloClient({
      uri: "http://localhost:3000/graphql",
      cache: new InMemoryCache(),
    });
    client
      .query({
        query: gql`
          {
            getProperty {
              id
              name
              created_at
              updated_at
            }
          }
        `,
      })
      .then((res) => {
        setRowData(res.data.getProperty);
      })
      .finally(() => {
        setIsLoading(false);
      });
  }, []);

  return (
    <div style={{ height: "88vh", width: "100%" }}>
      <CreateSection></CreateSection>
      <DeleteSection></DeleteSection>
      <DataGrid
        rows={rows}
        columns={columns}
        loading={isLoading}
        pagination
        pageSize={100}
        rowsPerPageOptions={[25, 50, 100, 250, 500, 1000]}
        // paginationMode='server'
        checkboxSelection
        disableSelectionOnClick
      />
    </div>
  );
};

export default Sections;
