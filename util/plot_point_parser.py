#!/usr/bin/env python3
"""Simple parser for the plot_point.txt table."""

from __future__ import annotations

from collections import Counter
from pathlib import Path
import re

# Basic patterns to keep the categorization simple and explainable.
ROLL_RANGE_RE = re.compile(r"^[0-9\s-]+$")
DEFINITION_RE = re.compile(r"^[A-Z0-9 ,!'&()\-]+:\s?.*")
ROLL_RANGE_FIELDS = ("Action", "Tension", "Mystery", "Social", "Personal")


def categorize_line(line: str) -> str:
    """Return a simple category label for the given line."""
    stripped = line.strip()
    if not stripped:
        return "blank"
    if ROLL_RANGE_RE.match(stripped):
        return "roll_range"
    if stripped.startswith("-"):
        return "table_entry"
    if stripped.isupper() and ":" not in stripped:
        return "heading"
    if DEFINITION_RE.match(stripped):
        return "definition"
    return "text"


def parse_roll_range_token(token: str) -> int:
    """Return the highest value represented by the token."""
    token = token.strip()
    if not token or token == "-":
        return 0
    numbers = [int(match) for match in re.findall(r"\d+", token)]
    return max(numbers, default=0)


def parse_roll_range_line(line: str) -> dict[str, int]:
    """Parse a roll range line into named numeric values."""
    tokens = line.split()
    values: dict[str, int] = {}
    for field, token in zip(ROLL_RANGE_FIELDS, tokens):
        values[field] = parse_roll_range_token(token)
    for field in ROLL_RANGE_FIELDS[len(tokens):]:
        values[field] = 0
    return values


def main() -> None:
    plot_point_path = Path(__file__).resolve().parent / "plot" / "plot_point.txt"
    lines = plot_point_path.read_text(encoding="utf-8").splitlines()

    totals: Counter[str] = Counter()
    results: list[dict[str, object]] = []

    line_index = 0
    while line_index < len(lines):
        line = lines[line_index]
        category = categorize_line(line)
        entry: dict[str, object] = {
            "index": line_index + 1,
            "type": category,
            "line": line,
        }

        if category == "roll_range":
            entry["ranges"] = parse_roll_range_line(line)
            results.append(entry)
            totals[category] += 1
            line_index += 1
            continue

        if category == "definition":
            text_lines: list[str] = []
            next_index = line_index + 1
            while next_index < len(lines):
                next_line = lines[next_index]
                next_category = categorize_line(next_line)
                if next_category == "text":
                    text_lines.append(next_line)
                    next_index += 1
                    continue
                if next_category == "roll_range":
                    break
                break
            if text_lines:
                entry["line"] = " ".join([line] + text_lines)
                entry["appended_text"] = text_lines
                totals["text"] += len(text_lines)
            results.append(entry)
            totals[category] += 1
            line_index = next_index
            continue

        results.append(entry)
        totals[category] += 1
        line_index += 1

    for result in results:
        base = f"{result['index']:04d} [{result['type']:11}] {result['line']}"
        if result["type"] == "roll_range":
            ranges = result.get("ranges", {})
            range_str = ", ".join(
                f"{field}={int(ranges.get(field, 0))}" for field in ROLL_RANGE_FIELDS
            )
            base = f"{base} -> {range_str}"
        print(base)

    print("\nCategory counts:")
    for category, count in totals.most_common():
        print(f"  {category:11}: {count}")


if __name__ == "__main__":
    main()
