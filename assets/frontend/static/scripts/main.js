document.addEventListener("DOMContentLoaded", () => {
  const shortenBtn = document.getElementById("shortenBtn");
  const urlInput = document.getElementById("urlInput");
  const resultDiv = document.getElementById("result");

  const listBtn = document.getElementById("listBtn");
  const passInput = document.getElementById("passInput");
  const linksTableContainer = document.getElementById("linksTableContainer");

  if (shortenBtn) {
    shortenBtn.addEventListener("click", async () => {
      const url = urlInput.value.trim();
      if (!url) return;

      try {
        const response = await fetch("/shortner", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ url })
        });
        
        resultDiv.style.display = "block";
        resultDiv.classList.remove("error");
        
        if (response.ok) {
          const alias = await response.text();
          const shortUrl = window.location.origin + "/redirect?alias=" + encodeURIComponent(alias);
          resultDiv.innerHTML = `URL shortened successfully!<br><br>Your short link:<br><a href="${shortUrl}" target="_blank" style="color:var(--accent-hover);font-weight:bold;word-break:break-all;">${shortUrl}</a>`;
        } else {
          resultDiv.classList.add("error");
          resultDiv.textContent = "Error: " + await response.text();
        }
      } catch (e) {
        resultDiv.style.display = "block";
        resultDiv.classList.add("error");
        resultDiv.textContent = "Network error occurred.";
      }
    });
  }

  if (listBtn) {
    listBtn.addEventListener("click", async () => {
      const password = passInput.value.trim();
      try {
        const response = await fetch(`/list-links?password=${encodeURIComponent(password)}`);
        if (response.ok) {
          const links = await response.json();
          renderLinksTable(links);
        } else {
          linksTableContainer.innerHTML = `<div class="result error" style="display:block">Error: ${await response.text()}</div>`;
        }
      } catch (e) {
        linksTableContainer.innerHTML = `<div class="result error" style="display:block">Network error occurred.</div>`;
      }
    });
  }

  function renderLinksTable(links) {
    if (!links || links.length === 0) {
      linksTableContainer.innerHTML = `<p style="text-align:center; color:var(--text-muted);">No links found.</p>`;
      return;
    }
    
    let html = `
      <table class="links-table fade-in">
        <thead>
          <tr>
            <th>Alias</th>
            <th>Original Link</th>
            <th>Redirects</th>
          </tr>
        </thead>
        <tbody>
    `;
    
    links.forEach(link => {
      html += `
        <tr>
          <td><a href="/redirect?alias=${encodeURIComponent(link.Alias)}" target="_blank">${link.Alias}</a></td>
          <td><a href="${link.Link}" target="_blank">${link.Link.substring(0, 50)}${link.Link.length > 50 ? '...' : ''}</a></td>
          <td>${link.Redirects}</td>
        </tr>
      `;
    });
    
    html += `</tbody></table>`;
    linksTableContainer.innerHTML = html;
  }
});
