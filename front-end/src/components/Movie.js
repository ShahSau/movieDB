import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

const Movie = () => {
    const [movie, setMovie] = useState({});
    let { id } = useParams();

    useEffect(() => {
        let myMovie ={
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
            }
        setMovie(myMovie);
    }, [id])

    return(
        <div>
            <h2>Movie: {movie.title}</h2>
            <small><em>{movie.release_date}</em></small>
            <hr />
            <p>{movie.description}</p>
        </div>
    )
}

export default Movie;