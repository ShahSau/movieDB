import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

const Movies = () => {
    const [movies, setMovies] = useState([]);

    useEffect( () => {
        let moviesList = [
            {
                adult: false,
                backdrop_path: "/r9oTasGQofvkQY5vlUXglneF64Z.jpg",
                id: 1029575,
                title: "The Family Plan",
                original_language: "en",
                original_title: "The Family Plan",
                overview: "Dan Morgan is many things: a devoted husband, a loving father, a celebrated car salesman. He's also a former assassin. And when his past catches up to his present, he's forced to take his unsuspecting family on a road trip unlike any other.",
                poster_path: "/jLLtx3nTRSLGPAKl4RoIv1FbEBr.jpg",
                media_type: "movie",
                genre_ids: [
                28,
                35
                ],
                popularity: 321.553,
                release_date: "2023-12-14",
                video: false,
                vote_average: 7.557,
                vote_count: 183
                },
                {
                    adult: false,
                    backdrop_path: "/aNG1BSIULxbHtUmBiar0i3fR1S4.jpg",
                    id: 520758,
                    title: "Chicken Run: Dawn of the Nugget",
                    original_language: "en",
                    original_title: "Chicken Run: Dawn of the Nugget",
                    overview: "A band of fearless chickens flock together to save poultry-kind from an unsettling new threat: a nearby farm that's cooking up something suspicious.",
                    poster_path: "/exNtEY8QUuQh9e23wSQjkPxKIU3.jpg",
                    media_type: "movie",
                    genre_ids: [
                    16,
                    12,
                    35,
                    10751
                    ],
                    popularity: 383.198,
                    release_date: "2023-12-08",
                    video: false,
                    vote_average: 7.6,
                    vote_count: 125
                    },
        ];

        setMovies(moviesList)
    }, []);

    return(
        <div>
            <h2>Movies</h2>
            <hr />
            <table className="table table-striped table-hover">
                <thead>
                    <tr>
                        <th>Movie</th>
                        <th>Release Date</th>
                        <th>Rating</th>
                    </tr>
                </thead>
                <tbody>
                    {movies.map((m) => (
                        <tr key={m.id}>
                            <td>
                                <Link to={`/movies/${m.id}`}>
                                    {m.title}
                                </Link>
                            </td>
                            <td>{m.release_date}</td>
                        </tr>    
                    ))}
                </tbody>
            </table>
        </div>
    )
}

export default Movies;