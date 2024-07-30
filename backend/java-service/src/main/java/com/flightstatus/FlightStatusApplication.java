package com.flightstatus;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.*;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.util.*;

@SpringBootApplication
public class FlightStatusApplication {
    public static void main(String[] args) {
        SpringApplication.run(FlightStatusApplication.class, args);
    }
}

@RestController
class FlightStatusController {

    private static final Map<String, String> USERS = new HashMap<>();
    private static final String SECRET_KEY = "super-secret";
    private final PasswordEncoder passwordEncoder = new BCryptPasswordEncoder();

    static {
        USERS.put("admin", new BCryptPasswordEncoder().encode("password123"));
    }

    @PostMapping("/api/login")
    public Map<String, String> login(@RequestBody Map<String, String> credentials, HttpServletResponse response) {
        String username = credentials.get("username");
        String password = credentials.get("password");

        if (username == null || password == null || !passwordEncoder.matches(password, USERS.get(username))) {
            response.setStatus(HttpServletResponse.SC_UNAUTHORIZED);
            return Collections.emptyMap();
        }

        String token = Jwts.builder()
                .setSubject(username)
                .setExpiration(new Date(System.currentTimeMillis() + 600000))
                .signWith(SignatureAlgorithm.HS512, SECRET_KEY)
                .compact();

        return Collections.singletonMap("token", token);
    }

    @GetMapping("/api/flights")
    public List<Flight> getFlights(HttpServletRequest request, HttpServletResponse response) {
        String token = request.getHeader("Authorization").replace("Bearer ", "");

        try {
            Jwts.parser().setSigningKey(SECRET_KEY).parseClaimsJws(token);
        } catch (Exception e) {
            response.setStatus(HttpServletResponse.SC_UNAUTHORIZED);
            return Collections.emptyList();
        }

        return Arrays.asList(
                new Flight(1, "On Time"),
                new Flight(2, "Delayed")
        );
    }

    static class Flight {
        private int id;
        private String status;

        public Flight(int id, String status) {
            this.id = id;
            this.status = status;
        }

        public int getId() {
            return id;
        }

        public String getStatus() {
            return status;
        }
    }
}
