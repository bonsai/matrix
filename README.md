# bonsai/matrix

Go-based Matrix Core for the Bonsai ecosystem.

Matrix is the **computational projection of Ontology + Synapse**. It provides reusable vector, matrix, matching, and data-interface primitives for Rails, BigQuery/BQML, CLI, and MCP/agents.

## Architecture

```text
Ontology + Synapse
        |
        v
   Matrix Core (Go)
     /    |    \
    v     v     v
Postgres  BQ    MCP
  |       |      |
Rails    BQML  Agents
```

## Principles

- Core mathematics is pure Go and has no network/GCP dependency.
- BigQuery is an adapter/analytical backend, not the domain model.
- MCP is an adapter over Matrix Core, not the calculation engine.
- `oshareco-tana` and other applications should consume this library rather than duplicate matching mathematics.

## Current packages

- `vector` — sparse vectors, normalization, dot product, norm
- `match` — similarity, complementarity, shared values, weighted score

Planned packages:

- `ontology` — ontology-weighted vector expansion
- `synapse` — weighted graph primitives
- `bq` — BigQuery adapters and feature projections
- `mcp` — agent-facing tools

## Match model

```text
M = ws*S + wc*C + wv*V
```

where `S` is similarity, `C` is complementarity, and `V` is shared values.

## Example

```go
w := match.Weights{Similarity: 0.4, Complementarity: 0.3, SharedValues: 0.3}
r := match.Match(a, b, w)
```

## Roadmap

1. Stable mathematical contracts
2. Ontology expansion
3. Synapse traversal
4. BigQuery interface / BQML feature projection
5. MCP adapter
6. Integration from `oshareco-tana`
