import SwaggerUI from 'swagger-ui'
import 'swagger-ui/dist/swagger-ui.css';
const axios = require('axios').default;

async function initSwagger() {
    const resp = await axios.get('/api/v1/swagger/spec.json')
    const ui = SwaggerUI({
        spec: resp.data,
        dom_id: '#swagger',
    });

    ui.initOAuth({
        appName: "Trickle Swagger UI",
        clientId: 'implicit'
    });
}

initSwagger()