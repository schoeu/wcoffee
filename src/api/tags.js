import axios from 'axios'

axios.get('http://localhost:8910/api/tags').then(res => {
    console.log(res);
}).catch(err => {
    console.log(err);
});


export default {
    getTags: 123
}