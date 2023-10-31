import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Button, Icon } from "semantic-ui-react";

let endpoint = "http://localhost:9000"

class WatchList extends Component {
    constructor(props) {
        super(props);
        this.state = {
            anime: "",
            items: [],
        };
    }
    componentDidMount() {
        this.getAnime();
    }

    onChange = (e) => {
        this.setState({
            [e.target.name]: e.target.value,
        })
    }

    // Enviando ---- IMPORTANTE
    onSubmit = () => {
        let { anime } = this.state;
        if (anime) {
            axios.post("/api/anime",
                { anime },
                {
                    headers:
                        { "Content-Type": "application/x-www-form-urlencoded" },
                },
            ).then(response => {
                this.getAnime();
                this.setState({
                    anime: "",
                });
                console.log(response);
            })
        }
    }


    getAnime = () => {
        // promissed do get
        axios.get("/api/anime").then((response) => {
            if (response.data) {
                this.setState({
                    items: response.data.map(item => {
                        let color = "orange"
                        let style = {
                            wordWrap: "break-word",
                        };
                        // sucesso
                        if (item.status) {
                            color = "green";
                            style["textDecorationLine"] = "line-through";
                        }
                        return (
                            <Card key={item._id} color={color} fluid className="rough">
                                <Card.Content>
                                    <Card.Header textAlign="left">
                                        <div style={style}>
                                            {item.anime}
                                        </div>

                                    </Card.Header>

                                    <Card.Meta textAlign="right">
                                        <Icon
                                            name="check circle"
                                            color="violet"
                                            onClick={
                                                () => { this.updateAnime(item._id); }
                                            }
                                        />
                                        <span style={{ paddingRight: 10 }}>Desfazer</span>
                                        <Icon
                                            name="delete"
                                            color="red"
                                            onClick={
                                                () => { this.deleteAnime(item._id) }
                                            }
                                        />
                                        <span style={{ paddingRight: 10 }}>Apagar</span>
                                    </Card.Meta>

                                </Card.Content>
                            </Card>
                        );
                    }),
                });
            } else {
                this.setState({
                    items: [],
                });
            }
        });
    };

    updateAnime = (id) => {
        axios.put("/api/anime" + id, {
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },
        }).then((response) => {
            console.log(response);
            this.getAnime();
        });
    }

    undoAnime = (id) => {
        axios.put("/api/descompletarAnime" + id, {
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },
        }).then((response) => {
            console.log(response);
            this.getAnime();
        });
    };

    deleteAnime = (id) => {
        axios.put("/api/apagarAnime" + id, {
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },
        }).then((response) => {
            console.log(response);
            this.getAnime();
        });
    };

    render() {
        return (
            <div className="main-div">
                <div className="row">
                    <Header className="header">

                        Animes para assistir
                    </Header>
                </div>
                <div className="row">
                    <Form className="" onSubmit={this.onSubmit}>
                        <Input
                            className="input-anime"

                            type="text"
                            name="anime"
                            onChange={this.onChange}
                            value={this.state.anime}
                            fluid
                            placeholder="Adicionar anime"
                        />
                        {<Button 
                        className="btn-add"> Add </Button>}
                    </Form>
                </div>
                <div className="row">
                    <Card.Group>
                        {this.state.items}
                    </Card.Group>
                </div>

            </div>
        )
    }
}

export default WatchList;