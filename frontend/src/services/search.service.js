import axios from "axios";
// import authHeader from "./auth-header";
// const API_URL = "http://localhost:8080/api/test/";
class UserService {
    getPosts(query) {
        return axios.get("/search/posts?q=" + query + "&l=50");
    }
    // getUserBoard() {
    //     return axios.get(API_URL + "user", { headers: authHeader() });
    // }
    // getModeratorBoard() {
    //     return axios.get(API_URL + "mod", { headers: authHeader() });
    // }
    // getAdminBoard() {
    //     return axios.get(API_URL + "admin", { headers: authHeader() });
    // }
}
export default new UserService();
