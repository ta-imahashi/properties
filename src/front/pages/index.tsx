import type { NextPage } from "next";
import { ApolloClient, InMemoryCache, gql } from "@apollo/client";

const test = async () => {
  const client = new ApolloClient({
    uri: "http://localhost:3000/graphql",
    cache: new InMemoryCache(),
  });
  const res = await client.query({
    query: gql`
      {
        findSection(id: 1) {
          id
          name
          type
          property_id
          created_at
          updated_at
        }
      }
    `,
  });
  console.log(res.data);
};

const Home: NextPage = () => {
  return (
    <button type="button" onClick={test}>
      Home
    </button>
  );
};

export default Home;
